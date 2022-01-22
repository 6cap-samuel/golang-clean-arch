package sam.henhaochi.authservice.controllers;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import sam.henhaochi.authservice.controllers.mappers.LoginAccountRequestMapper;
import sam.henhaochi.authservice.controllers.requests.LoginRequest;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.usecases.in.LoginAccountInput;

import java.security.NoSuchAlgorithmException;

@RestController
@AllArgsConstructor
@RequestMapping("/login")
public class LoginController {

    final LoginAccountInput loginAccountInput;
    final LoginAccountRequestMapper loginAccountRequestMapper;

    private static final Logger logger = LoggerFactory.getLogger(LoginController.class);

    @PostMapping
    public ResponseEntity<Object> loginWithUsernameAndPassword(
            @RequestBody LoginRequest loginRequest
    ) throws NoSuchAlgorithmException {
        logger.info("POST: /login called");
        Account foundAccount = loginAccountInput.with(
                loginAccountRequestMapper.mapToAccountEntity(
                        loginRequest
                )
        );

        if (foundAccount == null) {
            return ResponseEntity.status(HttpStatus.FORBIDDEN).build();
        } else {
            return ResponseEntity.status(HttpStatus.NO_CONTENT).build();
        }
    }
}
