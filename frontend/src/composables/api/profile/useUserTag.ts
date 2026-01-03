import { useApi } from '@/api/api';
import { useNotificationStore } from '@/stores/notification';
import { UserTagDTO } from '@/types/backend/user';

export function useUserTag() {
  const { notify } = useNotificationStore();
  
  const updateUserTag = async (dto: UserTagDTO) => {
    try {
      const { data } = await useApi('/api/user/tag', {
        credentials: 'include'
      })
        .put(dto).json();
        
        if (data.value.message !== null) {
          notify({
            type: "success",
            title: "User tag updated successfully",
            content: data.value.message
          });
        }
        
    } catch (e: any) {
        console.error(e);
    }
  };

  return {
    updateUserTag
  }
}