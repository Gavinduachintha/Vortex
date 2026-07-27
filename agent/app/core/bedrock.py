from langchain_aws import ChatBedrock

def get_model():
    return ChatBedrock(
        model_id="anthropic.claude-3-haiku-20240307-v1:0",
        region_name="ap-northeast-2",
    )