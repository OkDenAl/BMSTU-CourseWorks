import numpy as np


def min_max_filter(img, size):
    m, n = img.shape

    img_new = np.zeros([m, n])
    len = int((size - 1) / 2)
    temp = list()
    for i in range(len, m - len):
        for j in range(len, n - len):
            for k in range(i - len, i + len + 1):
                for w in range(j - len, j + len + 1):
                    temp.append(img[k, w])

            value = min(temp)
            img_new[i, j] = value
            temp.clear()

    img_new1 = np.zeros([m, n])
    for i in range(len, m - len):
        for j in range(len, n - len):
            for k in range(i - len, i + len + 1):
                for w in range(j - len, j + len + 1):
                    temp.append(img_new[k, w])

            value = max(temp)
            img_new1[i, j] = value
            temp.clear()

    img_new1 = img_new1.astype(np.uint8)

    return img_new1
