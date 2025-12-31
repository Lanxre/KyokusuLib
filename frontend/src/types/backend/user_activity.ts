export interface ActivityMetadata {
    ranobe_title?: string;
    author?: string;
    name?: string;
    desc?: string;
    [key: string]: any;
}

export interface UserActivity {
    id: number;
    activity_type: string;
    metadata: ActivityMetadata;
    timestamp: string;
}

export interface GetActivityResponse {
    message: UserActivity[];
}