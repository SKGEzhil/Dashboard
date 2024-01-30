import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class ContinueAsGuest extends StatelessWidget {
  const ContinueAsGuest({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 60,
      width: double.infinity,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(12),
        color: const Color(0xffFE724C).withOpacity(0.7),
      ),
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          onTap: () {
            //TODO: handle this
          },
          borderRadius: BorderRadius.circular(12),
          child: Center(
            child: Text(
              "Continue as Guest",
              style: GoogleFonts.inter(
                fontSize: 20,
                fontWeight: FontWeight.w500,
                color: const Color(0xff454545),
              ),
            ),
          ),
        ),
      ),
    );
  }
}