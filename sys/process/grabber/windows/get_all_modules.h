#ifndef TWITCH_CONNECTOR_GET_ALL_MODULES_H
#define TWITCH_CONNECTOR_GET_ALL_MODULES_H

#include <windows.h>

typedef struct module {
    char *baseName;
    char *fileName;
    uintptr_t baseAddr;
} module_t;

module_t **getAllModules(DWORD pid, DWORD *modCount);

#endif //TWITCH_CONNECTOR_GET_ALL_MODULES_H
