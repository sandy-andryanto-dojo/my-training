package com.userauth.service;

import org.springframework.security.core.userdetails.UserDetailsService;

import com.userauth.model.User;
import com.userauth.web.dto.UserRegistrationDto;

public interface UserService extends UserDetailsService {
	
	User save(UserRegistrationDto registrationDto);

}
