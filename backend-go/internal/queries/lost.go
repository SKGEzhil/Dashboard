
//HAVE TO UPDATE QUERIES AFTER CHANGING THE SCHEMA OF LOST AND FOUND TABLES 

package lost_found

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"strings"
)

type LostItem struct {
	ID              int      `db:"id"`
	ItemName        string   `db:"item_name"`
	ItemDescription string   `db:"item_description"`
	UserID          int      `db:"user_id"`
	Images          []string `db:"images"` 
	CreatedAt       string   `db:"created_at"`
}

type LostItemWithUser struct {
	ID              int    `db:"id"`
	ItemName        string `db:"item_name"`
	ItemDescription string `db:"item_description"`
	UserID          int    `db:"user_id"`
	UserName        string `db:"username"` 
	CreatedAt       string `db:"created_at"`
}

type ImageURI struct {
	ItemID   int    `db:"item_id"`
	ImageURL string `db:"image_url"`
}

func Insert_in_lost_table(db *sqlx.DB, form_data map[string]interface{}, user_ID int) (map[string]interface{}, error) {
	// Query to insert the lost item in the database
	query := `
        INSERT INTO lost (item_name, item_description, user_id) 
        VALUES ($1, $2, $3) 
        RETURNING *
    `

	// Execute the query
	var result map[string]interface{}
	err := db.Get(&result, query, form_data["item_name"], form_data["item_description"], user_ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Insert_lost_images(db *sqlx.DB, image_paths []string, post_id int) error {
	// Query to insert the lost item images in the database

	query := `INSERT INTO lost_images (image_url, item_id) 
        VALUES ($1, $2)`

	// 	Execute the query
	for _, image_paths := range image_paths {
		_, err := db.Exec(query, image_paths, post_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func Get_all_lost_items(db *sqlx.DB) ([]LostItem, error) {
	// Query to get all the lost items from the database
	query := `
        SELECT
            f.id,
            f.item_name,
            f.item_description,
            f.user_id,
            COALESCE(json_agg(fi.image_url) FILTER (WHERE fi.image_url IS NOT NULL), '[]') AS images,
            f.created_at
        FROM
            lost f
        LEFT JOIN
            lost_images fi ON f.id = fi.item_id
        GROUP BY
            f.id, f.item_name, f.item_description, f.user_id
        ORDER BY
            f.created_at DESC;
	`
	// Execute the query
	var lostItems []LostItem

	err := db.Select(&lostItems, query)
	if err != nil {
		return nil, err
	}
	return lostItems, nil
}

func Update_in_lost_table(db *sqlx.DB, item_id int, form_data map[string]interface{}) (LostItem, error) {
	// Query to update the lost item in the database

	query := `
		UPDATE lost SET
	`
	//List to hold the SET clause of the query
	setParts := []string{}

	// Arguments to pass to the query
	args := []interface{}{}
	argID := 1

	// Execute the query

	for key, value := range form_data {
		if key == "item_name" || key == "item_description" {
			setParts = append(setParts, fmt.Sprintf("%s = $%d", key, argID))
			args = append(args, value)
			argID++
		}
	}

	// Join the parts of the SET clause
	query += strings.Join(setParts, ", ")

	// Add the WHERE clause to filter by item ID
	query += fmt.Sprintf(" WHERE id = $%d", argID)
	args = append(args, item_id)

	// Add RETURNING * to get the updated row
	query += " RETURNING *"

	var updatedItem LostItem

	// Execute the query and scan the result into updatedItem
	err := db.QueryRowx(query, args...).StructScan(&updatedItem)
	if err != nil {
		return updatedItem, err
	}

	return updatedItem, nil

}

func Get_particular_lost_item(db *sqlx.DB, item_id int) (LostItemWithUser, error) {
	// Query to get the particular lost item from the database
	query := `
  SELECT 
	  f.id,
	  f.item_name,
	  f.item_description,
	  f.user_id,
	  u.username,
	  f.created_at
  FROM
	  lost f
  JOIN
	  users u ON f.user_id = u.id
  WHERE
	  f.id = $1
 `

	// Execute the query

	var lostItem LostItemWithUser
	err := db.Get(&lostItem, query, item_id)
	if err != nil {
		return lostItem, err
	}
	return lostItem, nil
}

func Delete_an_item_images_from_lost(db *sqlx.DB, item_id int) (string, error) {
	// Query to delete the particular lost item from the database
	query := `
	DELETE FROM lost_images
	WHERE item_id = $1
    `

	// Execute the query, passing the itemID as a parameter
	result_1, err := db.Exec(query, item_id)
	if err != nil {
		return "0", err
	}

	// Get the number of rows affected (i.e., number of images deleted)
	rowsAffected, err := result_1.RowsAffected()
	if err != nil {
		return "0", err
	}


	result := strings.Join([]string{fmt.Sprintf("%d", rowsAffected), " images deleted"}, "")

	return result, nil
}

func Delete_all_image_uris_lost(db sqlx.DB, item_id int) ([]string, error) {
	// Query to delete the particular lost item from the database
	query := `
    SELECT image_url FROM lost_images WHERE item_id = $1 
  `

	var image_urls []string
	err := db.Select(&image_urls, query, item_id)
	if err != nil {
		return nil, err
	}

	return image_urls, nil
}

func Search_lost_items_lost(db *sqlx.DB, search_query string) ([]LostItem, error) {
	max_results := 10
	// Query to search for lost items
	query := `
	  SELECT *
	  FROM lost
	  WHERE item_name ILIKE $1
	  ORDER BY created_at DESC
	  LIMIT $2
	`

	// Execute the query
	var lostItems []LostItem
	err := db.Select(&lostItems, query, "%"+search_query+"%", max_results)
	if err != nil {
		return nil, err
	}
	return lostItems, nil
}

func Get_some_img_uris_lost(db *sqlx.DB, item_ids []int) ([]ImageURI, error) {
	// Prepare a placeholder string for the query
	query, args, err := sqlx.In(`
   SELECT item_id, image_url
   FROM lost_images
   WHERE item_id IN (?)
`, item_ids)

	if err != nil {
		return nil, err
	}
	var image_uris []ImageURI
	query = db.Rebind(query)
	err = db.Select(&image_uris, query, args...)
	if err != nil {
		return nil, err
	}
	return image_uris, nil
}
