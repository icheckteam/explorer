import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing.unit * 3,
    overflowX: 'auto',
  },
  table: {
    minWidth: 700,
  },
});

function Transactions(props) {
  const { classes, transactions } = props;
  return (
    <Paper className={classes.root}>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell>Hash</TableCell>
            <TableCell>Height</TableCell>
            <TableCell>Type</TableCell>
            <TableCell>Fee</TableCell>
            <TableCell>Time</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {transactions.map(n => {
            return (
              <TableRow key={n.hash}>
                <TableCell>{n.hash}</TableCell>
                <TableCell>{n.height}</TableCell>
                <TableCell>{n.type}</TableCell>
                <TableCell>{n.fee}</TableCell>
                <TableCell>{moment(n.time).fromNow()}</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </Paper>
  );
}

Transactions.propTypes = {
  classes: PropTypes.object.isRequired,
  blocks: PropTypes.array
};

export default withStyles(styles)(Transactions);
