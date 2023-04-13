import os
import easyocr
import cv2
from matplotlib import pyplot as plt
import numpy as np
from transformers import (
  T5Tokenizer,
  MT5ForConditionalGeneration,
  Text2TextGenerationPipeline,
)


path = "K024/mt5-zh-ja-en-trimmed"
pipe = Text2TextGenerationPipeline(
  model=MT5ForConditionalGeneration.from_pretrained(path),
  tokenizer=T5Tokenizer.from_pretrained(path),
)
IMAGE_PATH = 'ja.jpg'
reader=easyocr.Reader(['ja','en'],gpu="False")
result=reader.readtext(IMAGE_PATH)
print(result)

for i in result:
    word=i[1]
    print(word)
    sentence="ja2zh: "+ word
    res = pipe(sentence, max_length=100, num_beams=4)
    print(res)
