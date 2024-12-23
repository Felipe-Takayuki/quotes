import 'package:diary_quoteapp/src/page/home/home_page.dart';
import 'package:flutter/material.dart';

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(home: Scaffold(
        body: DiaryQuote(),
    ),);
  }
}