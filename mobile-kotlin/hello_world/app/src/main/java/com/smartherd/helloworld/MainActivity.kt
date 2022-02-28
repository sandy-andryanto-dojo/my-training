package com.smartherd.helloworld

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.widget.Button
import android.widget.EditText
import android.widget.Toast

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val btnShowToast = findViewById<Button>(R.id.btnShowToast)
        val btnSendMsgToNextActivity = findViewById<Button>(R.id.btnSendMsgToNextActivity)
        val etUserMessage = findViewById<EditText>(R.id.etUserMessage)
        val btnShareToOtherApps = findViewById<Button>(R.id.btnShareToOtherApps)

        btnShowToast.setOnClickListener{
            Log.i("Main Activity", "Button was clicked !")

            Toast.makeText(this, "Button was clicked !", Toast.LENGTH_SHORT).show()

        }

        btnSendMsgToNextActivity.setOnClickListener{
            val message: String = etUserMessage.text.toString()
            var intent = Intent(this, SecondActivity::class.java)
            intent.putExtra("user_message", message)
            startActivity(intent)
        }

        btnShareToOtherApps.setOnClickListener{
            val message: String = etUserMessage.text.toString()
            var intent = Intent()
            intent.action = Intent.ACTION_SEND
            intent.putExtra(Intent.EXTRA_TEXT, message)
            intent.type = "text/plain"
            startActivity(Intent.createChooser(intent, "Please select app : "))
        }

    }
}