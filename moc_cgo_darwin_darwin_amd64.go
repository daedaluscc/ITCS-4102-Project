// +build !ios

package main

/*
#cgo CFLAGS: -pipe -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.9  -O2 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -stdlib=libc++ -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.9  -O2 -std=gnu++11 -Wall -W -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../GoglandProjects -I. -I../../Qt5.8.0/5.8/clang_64/lib/QtWidgets.framework/Headers -I../../Qt5.8.0/5.8/clang_64/lib/QtGui.framework/Headers -I../../Qt5.8.0/5.8/clang_64/lib/QtCore.framework/Headers -I. -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk/System/Library/Frameworks/OpenGL.framework/Headers -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk/System/Library/Frameworks/AGL.framework/Headers -I../../Qt5.8.0/5.8/clang_64/mkspecs/macx-clang -F/Users/joyrasmussen/Qt5.8.0/5.8/clang_64/lib
#cgo LDFLAGS: -headerpad_max_install_names -stdlib=libc++ -Wl,-syslibroot,/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk -mmacosx-version-min=10.9  -Wl,-rpath,/Users/joyrasmussen/Qt5.8.0/5.8/clang_64/lib
#cgo LDFLAGS:  -F/Users/joyrasmussen/Qt5.8.0/5.8/clang_64/lib -framework QtWidgets -framework QtGui -framework QtCore -framework DiskArbitration -framework IOKit -framework OpenGL -framework AGL
*/
import "C"
