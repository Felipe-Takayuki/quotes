class QuoteModel {
  String quote;
  String author;

  QuoteModel({required this.quote, required this.author});

  QuoteModel.fromJson(Map<String, dynamic> json, this.quote, this.author) {
    quote = json['quote'];
    author = json['author'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['quote'] = quote;
    data['author'] = author;
    return data;
  }
}