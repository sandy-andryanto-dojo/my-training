package com.sandy.latihan

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val btnCalculator = findViewById<Button>(R.id.btnKalkulator)
        btnCalculator.setOnClickListener{
            var intent = Intent(this, CalculatorActity::class.java)
            startActivity(intent)
        }


    }
}