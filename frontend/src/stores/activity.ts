import { defineStore, storeToRefs } from "pinia";
import { computed, watch } from "vue";
import { useIdle, useWindowFocus, useOnline, useIntervalFn } from "@vueuse/core";
import { useAuthStore } from "@/stores/auth";
import { useApi } from "@/api/api";
import { OnlineStatusEnum } from "@/types/enums/online-status-enum";

export const useActivityStore = defineStore("activity", () => {
    const authStore = useAuthStore();
    const { isAuthenticated } = storeToRefs(authStore);

    const { idle, lastActive } = useIdle(60 * 1000); 
    const focused = useWindowFocus();
    const online = useOnline();

    const isUserActive = computed(() => {
      return online.value && focused.value && !idle.value;
    });
    
    const sendHeartbeat = async () => {
        if (!isAuthenticated.value) return;

        try {
            await useApi("/api/user/activity", { credentials: "include" })
                .post({ 
                    status: isUserActive.value ? OnlineStatusEnum.ONLINE : OnlineStatusEnum.OFFLINE,
                    last_active: lastActive.value 
                });
        } catch (e) {
            console.error("Heartbeat failed", e);
        }
      };

    const { pause, resume } = useIntervalFn(() => {
        if (isUserActive.value) {
            sendHeartbeat();
        }
    }, 30000); 

    const initActivityTracking = () => {
        resume();
        
        watch(isAuthenticated, (isAuth) => {
            if (isAuth) {
                sendHeartbeat(); 
                resume();
            } else {
                pause();
            }
        }, { immediate: true });

        watch(isUserActive, (active) => {
            if (active && isAuthenticated.value) {
                sendHeartbeat();
            }
        });
    };

    return {
        isUserActive,
        initActivityTracking
    };
});