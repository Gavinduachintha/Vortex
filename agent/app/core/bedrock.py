from langchain_aws import ChatBedrock, BedrockEmbeddings


def get_chat_model():
    return ChatBedrock(
        model_id="anthropic.claude-3-haiku-20240307-v1:0",
        region_name="ap-northeast-2",
    )


def get_embedding_model():
    return BedrockEmbeddings(
        model_id="amazon.titan-embed-text-v2:0",
        region_name="ap-northeast-2",
    )