# Load required packages
import numpy as np
import time

# Load array data
arr = np.loadtxt(open("data/full_matrix.csv", "rb"), delimiter=",", skiprows=1)

# Calcualte loop start time
start_time = time.time()

# Loop through array
for i in range(len(arr)):
    for j in range(len(arr[0])):
        print(arr[i][j])

# Calculate total time to run loop 
print(f"--- {(round(time.time() - start_time, 3))} seconds to complete array loop ---")
