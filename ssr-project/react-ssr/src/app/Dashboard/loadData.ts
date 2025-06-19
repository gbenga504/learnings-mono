import { ICreateApiClient } from "../../api/api";

export const loadData = async ({
  api,
}: {
  api: ICreateApiClient;
}): Promise<{ id: string; name: string }[]> => {
  const result = await api.dashboard.getAllData();

  return result.data;
};
