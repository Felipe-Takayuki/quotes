import 'package:diary_quoteapp/src/models/quote_model.dart';
import 'package:diary_quoteapp/src/service/quote_service.dart';
import 'package:flutter/material.dart';

class HomeController extends ChangeNotifier {
    QuoteModel _quoteModel = QuoteModel(quote: "", author: "");

    QuoteModel get quote => _quoteModel;

    void loadDiaryQuote() async {
      _quoteModel = await QuoteService().getDiaryQuote();
      notifyListeners();
    }
}