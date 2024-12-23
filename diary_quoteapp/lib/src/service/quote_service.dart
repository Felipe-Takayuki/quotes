import 'package:diary_quoteapp/src/models/quote_model.dart';
import 'package:dio/dio.dart';

class QuoteService{
  Future<QuoteModel> getDiaryQuote() async {
    try{
      final dio =  Dio();
      final response = await dio.get('http://localhost:3000/quote');
      if (response.statusCode == 200) {
        final quote = QuoteModel(quote: response.data["quote"], author: response.data["author"]);
        return quote;
      }
    } catch (e) {
      Exception(e);
    }

    return QuoteModel(quote: "", author: "");
      
  }
}