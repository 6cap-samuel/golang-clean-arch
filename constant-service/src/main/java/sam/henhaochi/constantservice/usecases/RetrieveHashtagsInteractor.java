package sam.henhaochi.constantservice.usecases;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.usecases.in.RetrieveHashtagsInput;
import sam.henhaochi.constantservice.usecases.out.HashtagDataSource;

import java.util.List;

@AllArgsConstructor
@Service
public class RetrieveHashtagsInteractor
        implements RetrieveHashtagsInput {

    final HashtagDataSource hashtagDataSource;

    @Override
    public List<Hashtag> get() {
        return hashtagDataSource.retrieveAll();
    }
}
