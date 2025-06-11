import React, { useEffect, useState } from "react";
import {
  AppBar,
  Toolbar,
  Typography,
  Tabs,
  Tab,
  Box,
  Card,
  CardContent,
  CardHeader,
  Button,
  Avatar,
  Stack,
  IconButton,
  Collapse,
  List,
  ListItem,
  ListItemText,
  ListItemIcon,
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import ExpandLessIcon from "@mui/icons-material/ExpandLess";

interface Series {
  id: number;
  name: string;
  description: string;
  parent_id: number;
}

interface User {
  name: string;
  avatar: string;
}

function TabPanel(props: { children?: React.ReactNode; value: number; index: number }) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`tabpanel-${index}`}
      aria-labelledby={`tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

export default function BrickHolderApp() {
  const [user, setUser] = useState<User | null>(null);
  const [tabIndex, setTabIndex] = useState(1); // выбираем "Серии" по умолчанию
  const [series, setSeries] = useState<Series[]>([]);
  const [expandedParents, setExpandedParents] = useState<Set<number>>(new Set());

  useEffect(() => {
    async function fetchSeries() {
      try {
        // Пример запроса к API для получения серий
        const res = await fetch("/api/series"); // подставь свой URL
        const json = await res.json();
        setSeries(json.data);
      } catch (e) {
        console.error("Ошибка загрузки серий", e);
      }
    }
    fetchSeries();
  }, []);

  const handleLogin = () => {
    setUser({ name: "Alex", avatar: "https://i.pravatar.cc/100" });
  };

  const handleLogout = () => {
    setUser(null);
    setTabIndex(0);
  };

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setTabIndex(newValue);
  };

  // Функция для переключения раскрытия дочерних элементов у родителя
  const toggleExpand = (parentId: number) => {
    setExpandedParents((prev) => {
      const newSet = new Set(prev);
      if (newSet.has(parentId)) newSet.delete(parentId);
      else newSet.add(parentId);
      return newSet;
    });
  };

  // Отфильтровываем родителей (parent_id === 0)
  const parentSeries = series.filter((s) => s.parent_id === 0);

  // Для каждого родителя находим дочерние серии
  const getChildren = (parentId: number) => {
    return series.filter((s) => s.parent_id === parentId);
  };

  return (
    <Box sx={{ minHeight: "100vh", bgcolor: "#f5f5f5" }}>
      <AppBar position="static" color="default" elevation={1}>
        <Toolbar sx={{ justifyContent: "space-between" }}>
          <Stack direction="row" alignItems="center" spacing={2}>
            <img
              src="/brickholder_logo.png"
              alt="BrickHolder"
              style={{ height: 40 }}
            />
            <Typography variant="h6" component="div">
              BrickHolder
            </Typography>
          </Stack>

          <Stack direction="row" alignItems="center" spacing={2}>
            {user ? (
              <>
                <Avatar alt={user.name} src={user.avatar} />
                <Typography>{user.name}</Typography>
                <Button variant="outlined" onClick={handleLogout}>
                  Logout
                </Button>
              </>
            ) : (
              <Button variant="contained" onClick={handleLogin}>
                Login
              </Button>
            )}
          </Stack>
        </Toolbar>
      </AppBar>

      <Box sx={{ width: "100%", bgcolor: "background.paper" }}>
        <Tabs
          value={tabIndex}
          onChange={handleTabChange}
          aria-label="main tabs"
          variant="scrollable"
          scrollButtons="auto"
          sx={{ borderBottom: 1, borderColor: "divider" }}
        >
          <Tab label="Наборы" />
          <Tab label="Серии" />
          <Tab label="Детали" />
          {user && <Tab label="Моя коллекция" />}
          {user && <Tab label="Мои списки" />}
          {user && <Tab label="Статистика" />}
        </Tabs>

        <TabPanel value={tabIndex} index={0}>
          <Card>
            <CardHeader title="Серии LEGO" />
            <CardContent>
              <List>
                {parentSeries.map((parent) => {
                  const children = getChildren(parent.id);
                  const isExpanded = expandedParents.has(parent.id);
                  return (
                    <React.Fragment key={parent.id}>
                      <ListItem disablePadding>
                        <ListItem
                          onClick={() => toggleExpand(parent.id)}
                          sx={{ bgcolor: "#eee", mb: 1 }}
                        >
                          <ListItemText primary={parent.name} />
                          {isExpanded ? <ExpandLessIcon /> : <ExpandMoreIcon />}
                        </ListItem>
                      </ListItem>

                      <Collapse in={isExpanded} timeout="auto" unmountOnExit>
                        <List component="div" disablePadding sx={{ pl: 4 }}>
                          {children.length === 0 ? (
                            <ListItem>
                              <ListItemText primary="Нет дочерних серий" />
                            </ListItem>
                          ) : (
                            children.map((child) => (
                              <ListItem key={child.id} sx={{ pl: 2 }}>
                                <ListItemText primary={child.name} />
                              </ListItem>
                            ))
                          )}
                        </List>
                      </Collapse>
                    </React.Fragment>
                  );
                })}
              </List>
            </CardContent>
          </Card>
        </TabPanel>

        {/* Другие вкладки оставь как есть, например, Наборы, Детали и т.д. */}
      </Box>
    </Box>
  );
}
