'use strict';
const MANIFEST = 'flutter-app-manifest';
const TEMP = 'flutter-temp-cache';
const CACHE_NAME = 'flutter-app-cache';
const RESOURCES = {
  "assets/AssetManifest.json": "49782c420d760421033ff3d8b9e63466",
"assets/assets/fonts/Gloock-Regular.ttf": "8600463ccb102df8e60dbacd0384f94c",
"assets/assets/fonts/GreatVibes-Regular.ttf": "4bd0ed00ca8deac6c6706a092d9afd13",
"assets/assets/fonts/Lato-Black.ttf": "d83ab24f5cf2be8b7a9873dd64f6060a",
"assets/assets/fonts/Lato-Bold.ttf": "24b516c266d7341c954cb2918f1c8f38",
"assets/assets/fonts/Lato-Light.ttf": "2bcc211c05fc425a57b2767a4cdcf174",
"assets/assets/fonts/Lato-Regular.ttf": "122dd68d69fe9587e062d20d9ff5de2a",
"assets/assets/fonts/Lato-Thin.ttf": "7ab0bc06eecc1b75f8708aba3d3b044a",
"assets/assets/imgs/h1.png": "567ca513ae8c674e5448950bd6262aab",
"assets/assets/imgs/h2.png": "2d34cccee247e1e2335b74cfa724d596",
"assets/assets/imgs/h3.png": "756e1d95353a20651633a96658aeb6c6",
"assets/assets/imgs/h4.png": "724face9b7ab927fc2ed66cfb4339660",
"assets/assets/imgs/l1.png": "1dd1cb4e94b95f4fae377f91c3bbcb5b",
"assets/assets/imgs/p1.jpg": "ba2dd31626429f7d542ab81d07216730",
"assets/assets/imgs/s1.png": "c0f0b2c9b1178ef7508d670602d609c1",
"assets/assets/imgs/s2.png": "8c17750bd54539d76d75df3c50c61deb",
"assets/assets/imgs/s3.png": "c6671321ea675cd7b86d4251c06efd8c",
"assets/assets/imgs/s4.png": "4dca58753e550af71c59dd68ed6e7a29",
"assets/assets/imgs/s5.png": "9e68793650382a55aaf9f4f9f3b42a0b",
"assets/assets/imgs/s6.png": "e5c7bc583b0356aa7c924b34890c941e",
"assets/assets/imgs/s7.png": "53d2c261360f8d28a22efe7326e58abd",
"assets/assets/imgs/s8.png": "d0e5d08af080b699adf4a5dc0c9927cd",
"assets/FontManifest.json": "bb5296ba5563f2acf60f2c5af3ccbba6",
"assets/fonts/MaterialIcons-Regular.otf": "e7069dfd19b331be16bed984668fe080",
"assets/NOTICES": "e58f0144c0e362de1a803d62aecd4c99",
"assets/packages/cupertino_icons/assets/CupertinoIcons.ttf": "6d342eb68f170c97609e9da345464e5e",
"canvaskit/canvaskit.js": "97937cb4c2c2073c968525a3e08c86a3",
"canvaskit/canvaskit.wasm": "3de12d898ec208a5f31362cc00f09b9e",
"canvaskit/profiling/canvaskit.js": "c21852696bc1cc82e8894d851c01921a",
"canvaskit/profiling/canvaskit.wasm": "371bc4e204443b0d5e774d64a046eb99",
"favicon.png": "7e59b11568aae9e50f2b72da8c82d90e",
"flutter.js": "a85fcf6324d3c4d3ae3be1ae4931e9c5",
"icons/Icon-192.png": "32908504d2b2e058a8c29f6443204c57",
"icons/Icon-512.png": "0e241c794e13504103e16201a0537d8b",
"icons/Icon-maskable-192.png": "32908504d2b2e058a8c29f6443204c57",
"icons/Icon-maskable-512.png": "5f51e8f08cf5d20ccafefc93b46781a2",
"index.html": "e475894c2ee5a9a2161b7aabf1af226f",
"/": "e475894c2ee5a9a2161b7aabf1af226f",
"main.dart.js": "37485bd4921961c34c88a7cc3079c532",
"manifest.json": "ef14744cc0e4562c96495db5e01bc23a",
"version.json": "e72f9c18c1f0d0c5e6a8bdda53fd4ea5"
};

// The application shell files that are downloaded before a service worker can
// start.
const CORE = [
  "main.dart.js",
"index.html",
"assets/AssetManifest.json",
"assets/FontManifest.json"];
// During install, the TEMP cache is populated with the application shell files.
self.addEventListener("install", (event) => {
  self.skipWaiting();
  return event.waitUntil(
    caches.open(TEMP).then((cache) => {
      return cache.addAll(
        CORE.map((value) => new Request(value, {'cache': 'reload'})));
    })
  );
});

// During activate, the cache is populated with the temp files downloaded in
// install. If this service worker is upgrading from one with a saved
// MANIFEST, then use this to retain unchanged resource files.
self.addEventListener("activate", function(event) {
  return event.waitUntil(async function() {
    try {
      var contentCache = await caches.open(CACHE_NAME);
      var tempCache = await caches.open(TEMP);
      var manifestCache = await caches.open(MANIFEST);
      var manifest = await manifestCache.match('manifest');
      // When there is no prior manifest, clear the entire cache.
      if (!manifest) {
        await caches.delete(CACHE_NAME);
        contentCache = await caches.open(CACHE_NAME);
        for (var request of await tempCache.keys()) {
          var response = await tempCache.match(request);
          await contentCache.put(request, response);
        }
        await caches.delete(TEMP);
        // Save the manifest to make future upgrades efficient.
        await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
        return;
      }
      var oldManifest = await manifest.json();
      var origin = self.location.origin;
      for (var request of await contentCache.keys()) {
        var key = request.url.substring(origin.length + 1);
        if (key == "") {
          key = "/";
        }
        // If a resource from the old manifest is not in the new cache, or if
        // the MD5 sum has changed, delete it. Otherwise the resource is left
        // in the cache and can be reused by the new service worker.
        if (!RESOURCES[key] || RESOURCES[key] != oldManifest[key]) {
          await contentCache.delete(request);
        }
      }
      // Populate the cache with the app shell TEMP files, potentially overwriting
      // cache files preserved above.
      for (var request of await tempCache.keys()) {
        var response = await tempCache.match(request);
        await contentCache.put(request, response);
      }
      await caches.delete(TEMP);
      // Save the manifest to make future upgrades efficient.
      await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
      return;
    } catch (err) {
      // On an unhandled exception the state of the cache cannot be guaranteed.
      console.error('Failed to upgrade service worker: ' + err);
      await caches.delete(CACHE_NAME);
      await caches.delete(TEMP);
      await caches.delete(MANIFEST);
    }
  }());
});

// The fetch handler redirects requests for RESOURCE files to the service
// worker cache.
self.addEventListener("fetch", (event) => {
  if (event.request.method !== 'GET') {
    return;
  }
  var origin = self.location.origin;
  var key = event.request.url.substring(origin.length + 1);
  // Redirect URLs to the index.html
  if (key.indexOf('?v=') != -1) {
    key = key.split('?v=')[0];
  }
  if (event.request.url == origin || event.request.url.startsWith(origin + '/#') || key == '') {
    key = '/';
  }
  // If the URL is not the RESOURCE list then return to signal that the
  // browser should take over.
  if (!RESOURCES[key]) {
    return;
  }
  // If the URL is the index.html, perform an online-first request.
  if (key == '/') {
    return onlineFirst(event);
  }
  event.respondWith(caches.open(CACHE_NAME)
    .then((cache) =>  {
      return cache.match(event.request).then((response) => {
        // Either respond with the cached resource, or perform a fetch and
        // lazily populate the cache only if the resource was successfully fetched.
        return response || fetch(event.request).then((response) => {
          if (response && Boolean(response.ok)) {
            cache.put(event.request, response.clone());
          }
          return response;
        });
      })
    })
  );
});

self.addEventListener('message', (event) => {
  // SkipWaiting can be used to immediately activate a waiting service worker.
  // This will also require a page refresh triggered by the main worker.
  if (event.data === 'skipWaiting') {
    self.skipWaiting();
    return;
  }
  if (event.data === 'downloadOffline') {
    downloadOffline();
    return;
  }
});

// Download offline will check the RESOURCES for all files not in the cache
// and populate them.
async function downloadOffline() {
  var resources = [];
  var contentCache = await caches.open(CACHE_NAME);
  var currentContent = {};
  for (var request of await contentCache.keys()) {
    var key = request.url.substring(origin.length + 1);
    if (key == "") {
      key = "/";
    }
    currentContent[key] = true;
  }
  for (var resourceKey of Object.keys(RESOURCES)) {
    if (!currentContent[resourceKey]) {
      resources.push(resourceKey);
    }
  }
  return contentCache.addAll(resources);
}

// Attempt to download the resource online before falling back to
// the offline cache.
function onlineFirst(event) {
  return event.respondWith(
    fetch(event.request).then((response) => {
      return caches.open(CACHE_NAME).then((cache) => {
        cache.put(event.request, response.clone());
        return response;
      });
    }).catch((error) => {
      return caches.open(CACHE_NAME).then((cache) => {
        return cache.match(event.request).then((response) => {
          if (response != null) {
            return response;
          }
          throw error;
        });
      });
    })
  );
}
