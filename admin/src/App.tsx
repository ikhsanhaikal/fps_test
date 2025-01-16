import {
  Admin,
  Datagrid,
  DataProvider,
  DateField,
  fetchUtils,
  List,
  Resource,
  TextField,
} from "react-admin";

import queryString from "query-string";

function App() {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource name="produk" list={ListProduk} />
    </Admin>
  );
}

function ListProduk() {
  return (
    <List>
      <Datagrid>
        <TextField source="id" />
        <TextField source="nama_produk" />
        <TextField source="harga" />
        <TextField source="kategori_id" />
        <TextField source="status_id" />
        <DateField source="created_at" />
      </Datagrid>
    </List>
  );
}

const apiUrl = "http://localhost:3000";
const httpClient = fetchUtils.fetchJson;

const dataProvider: DataProvider = {
  getList: async (resource, params) => {
    let { page, perPage }: { page: number; perPage: number } =
      params.pagination;
    const url = `${apiUrl}/${resource}?${queryString.stringify({
      page,
      perPage,
    })}`;
    const resp = await httpClient(url, { signal: params.signal });
    return {
      data: resp.json.data,
      total: resp.json.total,
    };
  },
};

export default App;
