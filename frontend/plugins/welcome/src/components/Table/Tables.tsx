import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import moment from 'moment';
import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import { EntQueue } from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function ComponentsRecordQueueTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [queues, setQueues] = useState<EntQueue[]>([]);
  const [loading, setLoading] = useState(true);
  

  // get queues
  useEffect(() => {
    const getQueues = async () => {
      const res = await http.listQueue({ limit: 10, offset: 0 });
      setLoading(true);
      setQueues(res);
      console.log(res);
    };
    getQueues();
  }, [loading]);
  
  // delete button
  const deleteQueues = async (id: number) => {
    const res = await http.deleteQueue({ id: id });
    setLoading(true);
  };

    // clear input form
    function clear() {
      setQueues([]);
    }
  
 
  // ui 
 return (
  <Page theme={pageTheme.tool}>
    
    
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Patient</TableCell>
           <TableCell align="center">Dentist</TableCell>
           <TableCell align="center">Employee</TableCell>
           <TableCell align="center">Dental</TableCell>
           <TableCell align="center">Time</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {queues.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.patient?.patientName}</TableCell>
             <TableCell align="center">{item.edges?.dentist?.dentistName}</TableCell>
             <TableCell align="center">{item.edges?.employee?.employeeName}</TableCell>
             <TableCell align="center">{item.dental}</TableCell>
             <TableCell align="center">{moment(item.queueTime).format('DD/MM/YYYY HH:mm')}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deleteQueues(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href="/welcome"
               >
                 Delete
               </Button>
 
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}
 
