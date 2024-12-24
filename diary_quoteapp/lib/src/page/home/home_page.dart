
import 'package:diary_quoteapp/src/page/home/home_controller.dart';
import 'package:flutter/material.dart';

class DiaryQuote extends StatefulWidget {
  const DiaryQuote({super.key});

  @override
  State<DiaryQuote> createState() => _DiaryQuoteState();
}

class _DiaryQuoteState extends State<DiaryQuote> {
  
  HomeController quoteController = HomeController();
  @override
  void initState() {
    quoteController.loadDiaryQuote();
    super.initState();
  }
  @override
  Widget build(BuildContext context) {
      var sizeOf = MediaQuery.sizeOf(context);

    return  ListenableBuilder(listenable: quoteController, builder: (context,child) => 
    Padding(
      padding: const EdgeInsets.all(8.0),
      child: SizedBox(
        width: sizeOf.width * .98,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          spacing: 20,
          children: [
            Text("quote of the day:", style: TextStyle(fontWeight: FontWeight.bold, fontSize: 18),),
            
            quoteController.quote.quote.isEmpty ? 
            
            Padding(
              padding: const EdgeInsets.all(0),
              child: CircularProgressIndicator(),
            ) :
            
            Container(
              width: sizeOf.width *.75,
              padding: EdgeInsets.all(20),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(16),
                color: Color(0xffABA854)
              ),
              child: Center(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  spacing: 10,
                  children: [
                    Text("''${quoteController.quote.quote}''", style: TextStyle(fontSize: 26),),
                    Text(quoteController.quote.author, style: TextStyle(fontSize: 20, fontWeight: FontWeight.w700),)
                  ],
                ),
              ),
            ) 
          ],
        ),
      ),
    ));
  }
}