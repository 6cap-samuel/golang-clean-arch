package sam.henhaochi.authservice.annotations.methods;

import org.springframework.web.bind.annotation.CrossOrigin;

@CrossOrigin(origins = "${user.config.cors.url}")
public @interface WithCorsProtection { }