package sam.henhaochi.constantservice.controllers.responses;

import lombok.Builder;
import lombok.Data;
import sam.henhaochi.constantservice.controllers.structure.HashtagResponse;

import java.util.List;

@Builder
@Data
public class RetrieveHashtagsResponse {
    List<HashtagResponse> hashtags;
}
