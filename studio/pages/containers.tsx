import useSwr from 'swr'

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import { ChangeEvent } from 'react';
import { Button } from '@mui/material';

const fetcher = (url: string) => fetch(url).then((res) => res.json());
interface IResponse {
    code: number;
    data: any;
    message: string;
}

const Index = () => {

    const { data = [], mutate, error } = useSwr(`/favorite/containers`, fetcher);

    const handleChange = async function (e: ChangeEvent<HTMLInputElement>, checked: boolean)  {
        let operate = 'stop';
        if (checked) {
            // send stop
            operate = 'start';
        }
        const res = await fetch(`/favorite/containers/${operate}/${e.target.value}`).then<IResponse>((res):IResponse => res.json() as any);
        console.log(res);
        if (res.code === 200) {
            mutate(); // Refresh the data
        } else {
            alert(res.message);
        }
    }

    return (
        <div style={{padding: 20}}>
            <h1>Containers</h1>
            <TableContainer component={Paper} sx={{ maxHeight: '100%' }}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                    <TableRow>
                        <TableCell>ID</TableCell>
                        <TableCell align="right">Name</TableCell>
                        <TableCell align="right">State</TableCell>
                        <TableCell align="right">Action</TableCell>
                    </TableRow>
                    </TableHead>
                    <TableBody>
                    {data.map((row: any) => (
                        <TableRow
                        key={row.Id}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                        <TableCell component="th" scope="row">
                            {row.Id}
                        </TableCell>
                        <TableCell align="right">{row.Names}</TableCell>
                        <TableCell align="right">{row.State}</TableCell>
                        <TableCell align="right">
                            <FormControlLabel
                                label="启动"
                                control={<Switch
                                    value={row.Id}
                                    onChange={handleChange}
                                    checked={
                                    row.State === 'running'
                                }/>} 
                            />
                            <Button onClick={() => {
                                window.open(`http://localhost:${row.Ports[0].PublicPort}`);
                            }}>Preview</Button>
                        </TableCell>
                        </TableRow>
                    ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
}

export default Index;