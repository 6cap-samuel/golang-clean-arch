package sam.henhaochi.constantservice.repositories;

import lombok.AllArgsConstructor;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.stereotype.Repository;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.entities.HashtagEntity;
import sam.henhaochi.constantservice.repositories.mappers.HashtagEntityMapper;
import sam.henhaochi.constantservice.usecases.out.HashtagDataSource;

import java.util.List;
import java.util.stream.Collectors;

@Repository
@AllArgsConstructor
public class HashtagRepository implements HashtagDataSource {

    final MongoOperations mongoOperations;
    final HashtagEntityMapper mapper;

    @Override
    public List<Hashtag> retrieveAll() {
        return this.mongoOperations.findAll(
                HashtagEntity.class
        ).stream().map(mapper::map).collect(Collectors.toList());
    }
}
