package sam.henhaochi.constantservice.usecases.out;

import sam.henhaochi.constantservice.entities.Hashtag;

import java.util.List;

public interface HashtagDataSource {
    List<Hashtag> retrieveAll();
    void updateCount(List<String> hashtags);
}
