#include "get_all_modules.h"
#include <stdio.h>
#include <psapi.h>

module_t **getAllModules(DWORD pid, DWORD *modCount) {
    HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, pid);
    if (NULL == hProcess) {
        return NULL;
    }

    HMODULE hMods[2048];
    DWORD cbNeeded;

    if (!EnumProcessModulesEx(hProcess, hMods, sizeof(hMods), &cbNeeded, LIST_MODULES_ALL)) {
        CloseHandle(hProcess);
        return NULL;
    }

    *modCount = cbNeeded / sizeof(HMODULE);
    module_t **moduleList = calloc(1, sizeof(module_t));

    for (int i = 0; i < *modCount; i++) {
        if (i > 0) {
            moduleList = realloc(moduleList, (i + 1) * sizeof(module_t));
        }

        char baseName[MAX_PATH];
        char fileName[MAX_PATH];
        MODULEINFO modInfo;

        if (!GetModuleBaseNameA(hProcess, hMods[i], baseName, sizeof(baseName))) {
            CloseHandle(hProcess);
            return NULL;
        }

        if (!GetModuleFileNameExA(hProcess, hMods[i], fileName, sizeof(fileName))) {
            CloseHandle(hProcess);
            return NULL;
        }

        if (!GetModuleInformation(hProcess, hMods[i], &modInfo, sizeof(modInfo))) {
            CloseHandle(hProcess);
            return NULL;
        }

        moduleList[i] = calloc(1, sizeof(module_t));
        moduleList[i]->baseAddr = (uintptr_t) modInfo.lpBaseOfDll;

        moduleList[i]->baseName = calloc(1, strlen(baseName) + 1);
        strcpy(moduleList[i]->baseName, baseName);

        moduleList[i]->fileName = calloc(1, strlen(fileName) + 1);
        strcpy(moduleList[i]->fileName, fileName);
    }

    CloseHandle(hProcess);

    return moduleList;
}
