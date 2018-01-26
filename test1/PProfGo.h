#ifndef __PPROFGO_H__
#define __PPROFGO_H__

#ifdef _WIN32
#define DLLIMPORT __declspec(dllexport)
#else
#define DLLIMPORT
#endif

#ifdef __cplusplus
extern "C" {
#endif

    DLLIMPORT void _HeapProfilerStart(const char *prefix);
    DLLIMPORT void _HeapProfilerDump(const char *reason);
    DLLIMPORT void _HeapProfilerStop();

#ifdef __cplusplus
}
#endif

#endif
