// import 'package:flutter/material.dart';
// import 'package:dashbaord/screens/lost_and_found_add_item_screen.dart';
// import 'package:dashbaord/utils/normal_text.dart';

// class LostFoundAddItem extends StatelessWidget {
//   const LostFoundAddItem({
//     super.key,
//   });

//   @override
//   Widget build(BuildContext context) {
//     return Container(
//       clipBehavior: Clip.hardEdge,
//       decoration: BoxDecoration(
//           color: Theme.of(context).cardColor,
//           boxShadow: [
//             BoxShadow(
//               color: Color.fromRGBO(51, 51, 51, 0.10), // Shadow color
//               offset: Offset(0, 4), // Offset in the x, y direction
//               blurRadius: 10.0,
//               spreadRadius: 0.0,
//             ),
//           ],
//           borderRadius: BorderRadius.all(Radius.circular(10))),
//       child: Column(
//         crossAxisAlignment: CrossAxisAlignment.start,
//         children: [
//           Container(
//             color: Colors.white,
//             height: 150,
//             child: Center(
//               child: IconButton(
//                 onPressed: () => Navigator.of(context).push(
//                   MaterialPageRoute(
//                     builder: (context) => const LostAndFoundAddItemScreen(),
//                   ),
//                 ),
//                 icon: const Icon(
//                   Icons.add,
//                   size: 80,
//                   color: Color(0xB2FE724C),
//                 ),
//               ),
//             ),
//           ),
//           const SizedBox(height: 24),
//           const Padding(
//               padding: EdgeInsets.symmetric(horizontal: 10),
//               child: Center(
//                 child: NormalText(
//                   text: 'Add an Item',
//                   size: 16,
//                 ),
//               )),
//           const SizedBox(height: 10),
//         ],
//       ),
//     );
//   }
// }
