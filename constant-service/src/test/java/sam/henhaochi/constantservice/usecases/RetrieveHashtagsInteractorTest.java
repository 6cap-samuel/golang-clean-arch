package sam.henhaochi.constantservice.usecases;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.HashtagRepository;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.*;

class RetrieveHashtagsInteractorTest {

    List<Hashtag> results;

    private RetrieveHashtagsInteractor hashtagsInteractor;

    @BeforeEach
    void before(){
        results = new ArrayList<>();
        HashtagRepository hashtagRepository = mock(HashtagRepository.class);
        when(hashtagRepository.retrieveAll()).thenReturn(results);
        hashtagsInteractor = new RetrieveHashtagsInteractor(hashtagRepository);
    }

    @Test
    public void shouldCallHashtagDataSource(){
        assertEquals(
                results,
                hashtagsInteractor.get()
        );
    }
}