//go:build !customenv && staticx
// +build !customenv,staticx

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_aruco -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_superres -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_ximgproc -lopencv_imgcodecs -lopencv_flann -lopencv_xphoto -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_superres4100 -lopencv_gapi4100 -lopencv_aruco4100 -lopencv_ximgproc4100  -lopencv_videostab4100 -lopencv_video4100 -lopencv_rgbd4100 -lopencv_rapid4100 -lopencv_mcc4100 -lopencv_highgui4100 -lopencv_datasets4100 -lopencv_videoio4100 -lopencv_text4100 -lopencv_line_descriptor4100 -lopencv_imgcodecs4100 -lopencv_img_hash4100 -lopencv_hfs4100 -lopencv_fuzzy4100 -lopencv_xphoto4100 -lopencv_reg4100 -lopencv_quality4100 -lopencv_plot4100 -lopencv_photo4100 -lopencv_intensity_transform4100 -lopencv_imgproc4100 -lopencv_flann4100 -lopencv_core4100 -lade -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
