import { reactive, ref } from "vue";
import { useApi } from "@/api/api";
import { useNotificationStore } from "@/stores/notification";

export interface AuthorFormState {
    name: string;
    about: string;
    country: string;
    metier: string;
}

export function useAuthorForm() {
    const { notify } = useNotificationStore();

    const isLoading = ref(false);
    const isSuccess = ref(false);
    const errors = ref<Record<string, string>>({});
    
    const form = reactive<AuthorFormState>({
        name: "",
        about: "",
        country: "",
        metier: "",
    });

    const fileInput = ref<HTMLInputElement | null>(null);
    const selectedFile = ref<File | null>(null);
    const previewUrl = ref<string | null>(null);

    const validate = (): boolean => {
        errors.value = {};
        let isValid = true;

        if (!form.name.trim()) {
            errors.value.name = "Имя автора обязательно.";
            isValid = false;
        }

        if (!form.country.trim()) {
            errors.value.country = "Страна обязательна.";
            isValid = false;
        }
        
        if (!form.metier.trim()) {
            errors.value.metier = "Род деятельности обязателен.";
            isValid = false;
        }

        if (form.about.length > 2000) {
            errors.value.about = "Описание слишком длинное (макс. 2000 символов).";
            isValid = false;
        }

        return isValid;
    };

    const handleImageClick = () => {
        fileInput.value?.click();
    };

    const handleImageChange = (event: Event) => {
        const target = event.target as HTMLInputElement;
        if (target.files && target.files[0]) {
            const file = target.files[0];
            
            if (file.size > 5 * 1024 * 1024) {
                notify({
                    title: "Ошибка",
                    content: "Размер файла не должен превышать 5МБ",
                    type: "error"
                });
                return;
            }

            selectedFile.value = file;
            
            const reader = new FileReader();
            reader.onload = (e) => {
                if (e.target?.result) {
                    previewUrl.value = e.target.result as string;
                }
            };
            reader.readAsDataURL(file);
        }
    };

    const removeImage = () => {
        selectedFile.value = null;
        previewUrl.value = null;
        if (fileInput.value) fileInput.value.value = "";
    };

    const submit = async () => {
        if (!validate()) return;

        isLoading.value = true;
        isSuccess.value = false;

        try {
            const formData = new FormData();
            formData.append("name", form.name);
            formData.append("bio", form.about);
            formData.append("country", form.country);
            formData.append("metier", form.metier);
            
            if (selectedFile.value) {
                formData.append("picture", selectedFile.value); 
            }
            
            const { data, error  } = await useApi("/api/author", { credentials: "include" })
                .post(formData)
                .json();

            if (error.value) {
              if (typeof error.value === 'string') {
                  throw new Error(error.value);
              }
              
              const serverMessage = data.value?.error;
              throw new Error(serverMessage || "Неизвестная ошибка");
            }

            isSuccess.value = true;
            notify({
                title: "Успешно",
                content: "Автор успешно добавлен",
                type: "success"
            });

            form.name = "";
            form.about = "";
            form.metier = "";
            form.country = "";
            removeImage();

        } catch (e: any) {
            if (e.message) {
              errors.value.global = e.message;
              notify({
                  title: "Ошибка",
                  content: e.message,
                  type: "error",
              });
            } else {
                errors.value.global = "Неизвестная ошибка при создании автора";
                 notify({
                    title: "Ошибка",
                    content: "Неизвестная ошибка при создании автора",
                    type: "error",
                });
            }
        } finally {
            isLoading.value = false;
        }
    };

    return {
        form,
        errors,
        isLoading,
        isSuccess,
        fileInput,
        previewUrl,
        handleImageClick,
        handleImageChange,
        removeImage,
        submit
    };
}