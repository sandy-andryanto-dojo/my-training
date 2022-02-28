import 'package:flutter/material.dart';

void main()=> runApp(MaterialApp(
    home: Home()
));

class Home extends StatelessWidget{
    @override
    Widget build(BuildContext context){
      return Scaffold(
        appBar: AppBar(
          title: Text("Naon Atuh uyyyy"),
          centerTitle: true,
          backgroundColor: Colors.red[600],
        ),
        body: Center(
            child: Text(
              "This is Body",
              style: TextStyle(
                  fontSize: 20.0,
                  fontWeight: FontWeight.bold,
                  letterSpacing: 2.0,
                  color: Colors.grey[600]
              ),
            )
        ),
        floatingActionButton: FloatingActionButton(
          child: Text('click'),
          onPressed: (){},
          backgroundColor: Colors.red[600],
        ),
      );
    }
}
