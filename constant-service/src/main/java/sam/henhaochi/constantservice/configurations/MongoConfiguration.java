package sam.henhaochi.constantservice.configurations;

import com.mongodb.client.MongoClient;
import lombok.AllArgsConstructor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.MongoTemplate;

@AllArgsConstructor
@Configuration
public class MongoConfiguration {

    private static final String DB_NAME = "golang";

    final MongoClient mongoClient;

    @Bean
    public MongoOperations getMongoOperations() {
        return new MongoTemplate(
                mongoClient,
                DB_NAME
        );
    }
}
