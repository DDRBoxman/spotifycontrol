package spotifycontrol

/*
#cgo CFLAGS: -x objective-c -fobjc-arc 
#cgo LDFLAGS: -framework ScriptingBridge -framework Foundation
#import <Foundation/Foundation.h>
#import "Spotify.h"

SpotifyApplication*
spotify(void) {
	return [SBApplication applicationWithBundleIdentifier:@"com.spotify.client"];
}

void
nextTrack(void) {
	SpotifyApplication* app = spotify();
	[app nextTrack];	
}
*/
import "C"

func NextTrack() {
	C.nextTrack()
}
