import cv2
import numpy as np
import matplotlib as mpl
mpl.use('TkAgg')
from matplotlib import pyplot as plt
import json
from pprint import pprint

img = cv2.imread('img/jingcai_0.jpeg',0)
img2 = img.copy()

screenshot_h, screenshot_w = img.shape[::-1]

# All the 6 methods for comparison in a list
# methods = ['cv2.TM_CCOEFF', 'cv2.TM_CCOEFF_NORMED', 'cv2.TM_CCORR',
#             'cv2.TM_CCORR_NORMED', 'cv2.TM_SQDIFF', 'cv2.TM_SQDIFF_NORMED']

methods = ['cv2.TM_CCOEFF']

# Load the filename json
matched_char_count = 0
threshold = 0.8
team_left = []
team_right = []

with open('data/filename_simplified.json') as f:
    filenames = json.load(f)
    for fn in filenames:
        cname = fn['cname']
        filename = fn['filename']
        filename = 'img/' + filename + '.jpg'
        template = cv2.imread(filename,0)

        w, h = template.shape[::-1]

        for meth in methods:
            img = img2.copy()
            method = eval(meth)

            # Apply template Matching
            res = cv2.matchTemplate(img,template,method)
            min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(res)

            # If the method is TM_SQDIFF or TM_SQDIFF_NORMED, take minimum
            if method in [cv2.TM_SQDIFF, cv2.TM_SQDIFF_NORMED]:
                top_left = min_loc
            else:
                top_left = max_loc
            bottom_right = (top_left[0] + w, top_left[1] + h)

            print(cname + ': ')
            print(top_left)
            print(bottom_right)
            print(cv2.TM_CCOEFF)

            # cv2.rectangle(img,top_left, bottom_right, 255, 2)

            # plt.subplot(121),plt.imshow(res,cmap = 'gray')
            # plt.title('Matching Result'), plt.xticks([]), plt.yticks([])
            # plt.subplot(122),plt.imshow(img,cmap = 'gray')
            # plt.title('Detected Point'), plt.xticks([]), plt.yticks([])
            # plt.suptitle(meth)

            # plt.show()