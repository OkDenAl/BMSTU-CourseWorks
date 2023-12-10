import os
import time

import cv2
import numpy as np

from bilateral_filter import bilateral_filter
from mean_filter import mean_filter
from median_filter import median_filter
from min_max_filter import min_max_filter


def mean_filter_RGB(input_image):
    start = time.time()
    R_mf = mean_filter(input_image[:, :, 0])
    G_mf = mean_filter(input_image[:, :, 1])
    B_mf = mean_filter(input_image[:, :, 2])
    output_image_00 = np.stack([R_mf, G_mf, B_mf], axis=2)[:1080,:1080]
    mf_time = time.time() - start
    return output_image_00, mf_time


def bilateral_filter_RGB(input_image,kernel_size,alpha):
    start = time.time()
    R_bf = bilateral_filter(input_image[:, :, 0], kernel_size, alpha)
    G_bf = bilateral_filter(input_image[:, :, 1], kernel_size, alpha)
    B_bf = bilateral_filter(input_image[:, :, 2], kernel_size, alpha)
    output_image_01 = np.stack([R_bf, G_bf, B_bf], axis=2)
    bf_time = time.time() - start
    return output_image_01, bf_time


def median_filter_RGB(input_image, kernel_size):
    start = time.time()
    R_medf = median_filter(input_image[:, :, 0], kernel_size)
    G_medf = median_filter(input_image[:, :, 1], kernel_size)
    B_medf = median_filter(input_image[:, :, 2], kernel_size)
    output_image_02 = np.stack([R_medf, G_medf, B_medf], axis=2)
    medf_time = time.time() - start
    return output_image_02, medf_time


def min_max_filter_RGB(input_image,kernel_size):
    start = time.time()
    R_mmf = min_max_filter(input_image[:, :, 0], kernel_size)
    G_mmf = min_max_filter(input_image[:, :, 1], kernel_size)
    B_mmf = min_max_filter(input_image[:, :, 2], kernel_size)
    output_image_03 = np.stack([R_mmf, G_mmf, B_mmf], axis=2)
    mm_time = time.time() - start
    return output_image_03, mm_time

def test_algo():
    for i in range(1, 8):
        if not os.path.isdir(f'./output/{i}'):
            os.mkdir(f'./output/{i}')

        print(f'-----------------STARTING PROCESSING IMAGE {i}------------------')
        image_name = f'./dataset/{i}.png'
        input_image = cv2.imread(image_name,
                                 cv2.IMREAD_UNCHANGED).astype(np.float32) / 255.0

        # mean filter
        print("Starting mean filter...")
        output_image_00, mf_time = mean_filter_RGB(input_image)
        cv2.imwrite(f'./output/{i}/mean.png', output_image_00 * 255)
        print("Mean filter end")

        # bilateral
        print("Starting bilateral filter...")
        output_image_01, bf_time = bilateral_filter_RGB(input_image,5,0.6)
        cv2.imwrite(f'./output/{i}/bilateral.png', output_image_01 * 255)
        print("Bilateral filter end")

        # median
        print("Starting median filter...")
        output_image_02, medf_time = median_filter_RGB(input_image,5)
        cv2.imwrite(f'./output/{i}/median.png', output_image_02 * 255)
        print("Median filter end")

        # # min-max
        print("Starting min-max filter...")
        output_image_03, mm_time = min_max_filter_RGB(input_image,5)
        cv2.imwrite(f'./output/{i}/min_max.png', output_image_03 * 255)
        print("Min-max filter end")

        print(f'\n--------------------TIME-----------------------\n')
        print(f'{mf_time}\t{bf_time}\t{medf_time}\t{mm_time}')

        input_image = np.stack([input_image[:,:,0],input_image[:,:,1],input_image[:,:,2]],axis=2)

        Row1 = np.hstack([input_image,output_image_00, output_image_01,
                          output_image_02,output_image_03])

        # write out the image
        cv2.imwrite(f'./output/{i}_all.png', Row1*255)
