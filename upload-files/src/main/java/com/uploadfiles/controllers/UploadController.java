package com.uploadfiles.controllers;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

@Controller
public class UploadController {
	
	@Autowired
	private Environment env;
	
	@GetMapping("/")
	public String viewHomePage() {
		return "index";
	}
	
	@GetMapping("/uploadStatus")
    public String uploadStatus() {
        return "uploadStatus";
    }
	
	@PostMapping("/upload") // //new annotation since 4.3
    public String singleFileUpload(@RequestParam("file") MultipartFile file, RedirectAttributes redirectAttributes) {

        if (file.isEmpty()) {
            redirectAttributes.addFlashAttribute("message", "Please select a file to upload");
            return "redirect:uploadStatus";
        }
        
        
        try {
        	
        	UUID uuid = UUID.randomUUID();
        	String uploadPath = env.getProperty("upload.path");
        	String filenameReal = file.getOriginalFilename();
        	String fileExtension = filenameReal.substring(filenameReal.lastIndexOf(".") + 1);
        	String fileName = uuid+"."+fileExtension;
        	
            // Get the file and save it somewhere
            byte[] bytes = file.getBytes();
            Path path = Paths.get(uploadPath + fileName);
            Files.write(path, bytes);

            redirectAttributes.addFlashAttribute("message",
                    "You successfully uploaded '" + filenameReal + "'");

        } catch (IOException e) {
            e.printStackTrace();
        }

        return "redirect:/";
    }
	
	private String getFileExtension(File file) {
	    String name = file.getName();
	    int lastIndexOf = name.lastIndexOf(".");
	    if (lastIndexOf == -1) {
	        return ""; // empty extension
	    }
	    return name.substring(lastIndexOf);
	}

}
