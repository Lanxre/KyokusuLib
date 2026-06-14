export interface NovelaShort {
  novelaId: number;
  title: string;
  description: string;
  posterUrl: string;
  releaseDate: Date;
}

export interface TotalNovelaStatistics {
  novela: NovelaShort;
  bookmarkCount: number;
  readCount: number;
  ratingCount: number;
  commentCount: number;
  volumeCount: number;
  chapterCount: number;
  rating: number;
}

export interface TotalNovelaStatisticsResponse {
  data: TotalNovelaStatistics[];
  total: number;
}

export interface GeneralStatistics {
  bookmarkCount: number;
  readCount: number;
  ratingCount: number;
  commentCount: number;
  volumeCount: number;
  chapterCount: number;
}
