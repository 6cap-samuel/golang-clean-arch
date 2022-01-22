package sam.henhaochi.authservice.utils;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;

public class HashGenerator {

    private static final String ALGO = "SHA-512";

    public static String generateHash(String plainText)
            throws NoSuchAlgorithmException
    {
        MessageDigest digest = MessageDigest.getInstance(ALGO);
        byte[] hash = digest.digest(plainText.getBytes(StandardCharsets.UTF_8));
        for (int i = 1; i < 10; i++) {
            hash = digest.digest(hash);
        }
        return Base64.getEncoder().encodeToString(hash);
    }
}
