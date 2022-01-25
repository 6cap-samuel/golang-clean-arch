package sam.henhaochi.constantservice.repositories;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.data.mongodb.core.MongoOperations;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.entities.HashtagEntity;
import sam.henhaochi.constantservice.repositories.mappers.HashtagEntityMapper;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class HashtagRepositoryTest {

    private static String ID = "";
    private static String NAME = "";
    private static int COUNT = 0;

    List<HashtagEntity> results;

    HashtagRepository hashtagRepository;
    MongoOperations mongoOperations;
    HashtagEntityMapper mapper;

    @BeforeEach
    void before(){
        mongoOperations = mock(MongoOperations.class);
        mapper = mock(HashtagEntityMapper.class);

        hashtagRepository = new HashtagRepository(
                mongoOperations,
                mapper
        );
    }

    @Test
    public void shouldReturnListOfHashtag() {
        results = new ArrayList<>();
        HashtagEntity hashtagEntity =
                HashtagEntity.builder()
                        .id(ID)
                        .count(COUNT)
                        .name(NAME)
                        .build();
        results.add(hashtagEntity);

        Hashtag hashtag = Hashtag.builder()
                .id(ID)
                .count(COUNT)
                .name(NAME)
                .build();

        when(mongoOperations.findAll(
                HashtagEntity.class
        )).thenReturn(results);

        when(mapper.map(
                hashtagEntity
        )).thenReturn(
                hashtag
        );

        assertEquals(
                hashtag,
                hashtagRepository.retrieveAll().get(0)
        );
    }
}