import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import ComponanceTable from '../Table';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(70),
    },
  }),
);

const WelcomePage: FC<{}> = () => {
  //const classes = useStyles();
 return (
   
   <Page theme={pageTheme.app}>
     <Header
       title="Dental System"
       subtitle="ระบบจองคิวผู้ป่วย">
     </Header>

      
        <Content>
        <ContentHeader title="ตารางคิวผู้ป่วย">
        <Link component={RouterLink} to="/">
              <Button
                variant="contained"
                color="default"
              >
                ออกจากระบบ
              </Button>
              </Link>
              &emsp;
        <Link component={RouterLink} to="/queueorder">
            <Button variant="contained" color="primary">
              จองคิว
            </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
        </Content>
      
   </Page>
 );
};

export default WelcomePage;