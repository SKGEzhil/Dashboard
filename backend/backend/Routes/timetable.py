from fastapi import APIRouter, HTTPException, Response
from utils import conn
from models import Timetable
from queries import timetable as timetable_queries
from psycopg2.errors import UniqueViolation, ForeignKeyViolation
router = APIRouter(prefix="/timetable", tags=["timetable"])


@router.get("/{user_id}")
async def get_timetable(user_id: int, acad_period: str, res: Response) -> Timetable:
    try:
        query = timetable_queries.get_timeTable(user_id, acad_period)
        
        with conn.cursor() as cur: 
            cur.execute(query)

            rows = cur.fetchall()
            res.status_code = 200  
            
            timetable = Timetable(user_id=user_id, acad_period=acad_period, course_codes=[
                row[0] for row in rows
            ])
            
            return timetable
    except Exception as e:
        return HTTPException(status_code=500, detail=f"Internal Server Error : {e}")


@router.post("/") # the route also works to update the timetable 
async def update_timetable(timetable: Timetable, res: Response):
    
    try:
        user_id = timetable.user_id
        acad_period = timetable.acad_period
        query = timetable_queries.get_allRegisteredCourses(user_id, acad_period)
        
        with conn.cursor() as cur: # for transaction
            cur.execute(query)
            conn.commit()
            
            rows = cur.fetchall()
            courses = [row[0] for row in rows]


            for course_code in timetable.course_codes: 
                if course_code in courses:
                    timetable.course_codes.remove(course_code)
                    courses.remove(course_code)

            timetable.course_codes = list(set(timetable.course_codes)) # having unique courses
            
            if(len(timetable.course_codes) == 0 and len(courses) == 0): # if no courses are to be added or deleted
                res.status_code = 200
                return {"message" : "No Courses to update"}
            
            
            if(len(courses) != 0): # if some courses are to be deleted
                query = timetable_queries.delete_timeTable(Timetable(user_id=user_id, acad_period=acad_period, course_codes=courses))
                cur.execute(query)
                
            if(len(timetable.course_codes) != 0): # if some courses are to be added
                query = timetable_queries.post_timeTable(timetable)
                cur.execute(query)
                
            res.status_code = 200
            return {"message" : "Courses Updated Successfully"}
        
    except ForeignKeyViolation as e: # if some timetable doesn't exist
        return HTTPException(status_code=404, detail=f"Course not Found : {e}")
    except Exception as e:
        return HTTPException(status_code=500, detail=f"Internal Server Error : {e}")
    
