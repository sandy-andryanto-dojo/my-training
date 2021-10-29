package com.console.console;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class ConsoleApplication implements CommandLineRunner{

	public static void main(String[] args) {
        SpringApplication.run(ConsoleApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		for(int i = 1; i <=1000; i++) {
			System.out.println("Hello World "+i);
		}
	}
}
