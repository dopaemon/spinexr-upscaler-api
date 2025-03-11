import argparse
import shutil
import os
import pydicom
import numpy as np
from PIL import Image

def convert_dicom_to_png(dicom_path, png_path):
    dicom_data = pydicom.dcmread(dicom_path)
    img_array = dicom_data.pixel_array
    img = Image.fromarray((img_array / img_array.max() * 255).astype(np.uint8))
    img.save(png_path)

def process_image(input_path, output_path):
    os.makedirs(os.path.dirname(output_path), exist_ok=True)
    ext = os.path.splitext(input_path)[1].lower()
    if ext in [".dcm", ".dicom"]:
        dicom_output = output_path.replace(".png", ".dcm")
        shutil.copy(input_path, dicom_output)
        convert_dicom_to_png(input_path, output_path)
    else:
        shutil.copy(input_path, output_path)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Process image")
    parser.add_argument("-i", "--input", required=True, help="Input image path")
    parser.add_argument("-o", "--output", required=True, help="Output image path")
    args = parser.parse_args()
    process_image(args.input, args.output)
