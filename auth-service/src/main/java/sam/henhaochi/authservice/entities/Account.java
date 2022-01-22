package sam.henhaochi.authservice.entities;

import lombok.Builder;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import sam.henhaochi.authservice.utils.HashGenerator;

import java.security.NoSuchAlgorithmException;

@Data
@Builder
@Getter
@Setter
public class Account {

    private String username;
    private String email;
    private String password;

    public void encrypt() throws NoSuchAlgorithmException {
        password = HashGenerator.generateHash(password);
    }
}
