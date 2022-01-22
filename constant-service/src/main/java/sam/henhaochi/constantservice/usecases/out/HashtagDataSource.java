package sam.henhaochi.constantservice.usecases.out;

import sam.henhaochi.constantservice.entities.Hashtag;

import java.util.List;
import java.util.Set;

public interface HashtagDataSource {
    List<Hashtag> retrieveAll();
    void updateCount(List<String> hashtags);
}
