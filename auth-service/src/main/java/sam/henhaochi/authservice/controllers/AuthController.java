package sam.henhaochi.authservice.controllers;

import lombok.AllArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import sam.henhaochi.authservice.annotations.methods.WithCorsProtection;
import sam.henhaochi.authservice.controllers.mappers.LoginAccountRequestMapper;
import sam.henhaochi.authservice.controllers.mappers.LoginAccountResponseMapper;
import sam.henhaochi.authservice.controllers.requests.LoginRequest;
import sam.henhaochi.authservice.entities.Account;
import sam.henhaochi.authservice.usecases.in.LoginAccountInput;

import java.security.NoSuchAlgorithmException;

@RestController
@AllArgsConstructor
@RequestMapping("/auth")
public class AuthController {

    final LoginAccountInput loginAccountInput;
    final LoginAccountRequestMapper loginAccountRequestMapper;
    final LoginAccountResponseMapper loginAccountResponseMapper;

    private static final Logger logger = LoggerFactory.getLogger(AuthController.class);

    @PostMapping("/login")
    @WithCorsProtection
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
            return ResponseEntity.status(
                    HttpStatus.FORBIDDEN
            ).build();
        } else {
            return ResponseEntity.ok(
                    loginAccountResponseMapper.mapToResponseEntity(
                            foundAccount
                    )
            );
        }
    }
}
