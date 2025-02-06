# VODM - Simple File Downloader

Working on [3steps](https://3steps.no) We often need to analyze plenty of handball matches. Therefore, we need to download plenty of handball matches. This tool lets us download several matches at once and doesn't use up all the computer's resources so that we can still work while using it.

![Screenshot 2025-02-06 at 20 49 17](https://github.com/user-attachments/assets/12f3f8fe-a77a-4761-9a62-572ead9ea66f)


## Usage

```bash
vodm URL1 [URL2...]
# Example:
vodm https://example.com/file1.mp4 https://example.com/file2.mp4
```

4. Create a `.gitignore`:
```gitignore
# Binary
vodm

# macOS
.DS_Store
```

****

## Installation

### macOS (M1/Intel)

```bash
# Clone the repository
git clone https://github.com/yourusername/vodm.git
cd vodm

# Make install script executable
chmod +x bin/install.sh

# Run installation
./bin/install.sh
```

## Passing text file
```bash
vodm -c all_urls.txt
```

## Concurrent installation
Using the `-c` flag triggers concurrent downloads. you can increase and decrease the amount of workers by using the flag `-l 5`.

```bash
vodm -c https://videos.pexels.com/video-files/29538074/12714837_360_640_60fps.mp4 https://videos.pexels.com/video-files/30075744/12899829_640_360_30fps.mp4 https://videos.pexels.com/video-files/6788661/6788661-sd_640_360_25fps.mp4
```

## Output
Use the `-o /path` argument to output the files to other destinations
