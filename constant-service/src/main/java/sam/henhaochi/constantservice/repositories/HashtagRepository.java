package sam.henhaochi.constantservice.repositories;

import com.mongodb.client.MongoClient;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Repository;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.entities.HashtagEntity;
import sam.henhaochi.constantservice.repositories.mappers.HashtagEntityMapper;
import sam.henhaochi.constantservice.usecases.out.HashtagDataSource;

import java.util.List;
import java.util.stream.Collectors;

import static org.springframework.data.mongodb.core.query.Criteria.where;

@Repository
public class HashtagRepository implements HashtagDataSource {

    final MongoOperations mongoOperations;
    final HashtagEntityMapper mapper;

    public HashtagRepository(
            final MongoClient mongoClient,
            final HashtagEntityMapper hashtagEntityMapper
    ) {
        this.mongoOperations = new MongoTemplate(
                mongoClient,
                "golang"
        );
        this.mapper = hashtagEntityMapper;
    }

    public List<HashtagEntity> findAll(List<String> hashtagString){
        return this.mongoOperations.find(
                new Query(where("name").in(hashtagString)), HashtagEntity.class
        );
    }

    public void insertHashtags(List<String> hashtags){
        this.mongoOperations.insertAll(
                hashtags.stream().map(mapper::mapStringToEntity).collect(Collectors.toList())
        );
    }

    public void updateHashtags(List<HashtagEntity> hashtags) {
        hashtags.forEach(
                (a) -> this.mongoOperations.updateFirst(
                        new Query(Criteria.where("name").in(a.getName())),
                        new Update().set("count", a.incrementCount()),
                        HashtagEntity.class
                )
        );
    }

    @Override
    public void updateCount(
            List<String> hashtags
    ) {
        List<HashtagEntity> foundHashtag = findAll(hashtags);
        updateHashtags(foundHashtag);
        hashtags = hashtags.stream().filter(
                (a) -> foundHashtag.stream().filter(
                        (b) -> b.getName().equals(a)
                ).collect(Collectors.toSet()).size() != 0)
                .collect(Collectors.toList());
        insertHashtags(hashtags);
    }

    @Override
    public List<Hashtag> retrieveAll() {
        return this.mongoOperations.findAll(
                HashtagEntity.class
        ).stream().map(mapper::map).collect(Collectors.toList());
    }
}
