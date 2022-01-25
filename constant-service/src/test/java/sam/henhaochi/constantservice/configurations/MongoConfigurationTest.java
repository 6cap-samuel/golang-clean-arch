package sam.henhaochi.constantservice.configurations;

import com.mongodb.client.MongoClient;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.mock;

class MongoConfigurationTest {

    MongoClient mongoClient;
    MongoConfiguration mongoOperations;

    @BeforeEach
    void before(){
        mongoClient = mock(MongoClient.class);
        mongoOperations = new MongoConfiguration(
                mongoClient
        );
    }

    @Test
    public void shouldReturnNewMongoOperationsInstance() {
        assertNotNull(
                mongoOperations.getMongoOperations()
        );
    }
}