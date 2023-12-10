from skimage.metrics import structural_similarity
import cv2


def count_SSMI():
    filters = ["median", "bilateral", "mean"]
    for i in range(1, 8):
        for filter in filters:
            before = cv2.imread(f'./dataset/{i}.png')
            after = cv2.imread(f'./output/{i}/{filter}.png')

            before_gray = cv2.cvtColor(before, cv2.COLOR_BGR2GRAY)
            after_gray = cv2.cvtColor(after, cv2.COLOR_BGR2GRAY)

            (score, diff) = structural_similarity(before_gray, after_gray, full=True)
            print(f'image similarity source with {filter}', score)
