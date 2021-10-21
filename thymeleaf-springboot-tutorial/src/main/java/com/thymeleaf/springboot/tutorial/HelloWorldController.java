package com.thymeleaf.springboot.tutorial;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HelloWorldController {
	
	@GetMapping("/hello")
	public String hello(Model model) {
		model.addAttribute("message", "Hello World");
		return "helloworld";
	}

	@GetMapping("/hello_css")
	public String helloCss() {
		return "hello_css";
	}
	
	@GetMapping("/bootstrap")
	public String bootstrap() {
		return "bootstrap";
	}
}
