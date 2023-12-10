import argparse

import cv2
import numpy as np

import testing
from ssim import count_SSMI


def create_parser():
    pars = argparse.ArgumentParser()
    pars.add_argument('-m', type=str, default='test1', help="program mode")
    pars.add_argument('-f', type=str, default='median', help="filters")
    pars.add_argument('-ks', type=int, default=5, help="kernel size")
    pars.add_argument('-a', type=int, default=0.6, help="alpha (0,1)")
    pars.add_argument('-inp', type=str, default='./dataset/3.png', help="path to input image")
    pars.add_argument('-out', type=str, default='./out/3_done.png', help="path to output image")

    return pars


def main():
    parser = create_parser()
    console_args = parser.parse_args()
    if console_args.m == "test":
        testing.test_algo()
        count_SSMI()
    elif console_args.m == "prod":
        input_image = cv2.imread(console_args.inp, cv2.IMREAD_UNCHANGED).astype(np.float32) / 255.0
        filters = console_args.f.split()
        for filter in filters:
            if filter == "median":
                output_image, _ = testing.median_filter_RGB(input_image, console_args.ks)
                cv2.imwrite(console_args.out, output_image * 255)
            elif filter == "mean":
                output_image, _ = testing.mean_filter_RGB(input_image)
                cv2.imwrite(console_args.out, output_image * 255)
            elif filter == "bilateral":
                output_image, _ = testing.bilateral_filter_RGB(input_image, console_args.ks, console_args.alpha)
                cv2.imwrite(console_args.out, output_image * 255)
            else:
                print("cant use this filter")
    else:
        print("unknown mode")


main()
