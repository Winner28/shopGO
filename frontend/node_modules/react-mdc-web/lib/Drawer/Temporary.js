'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _DrawerShade = require('./DrawerShade');

var _DrawerShade2 = _interopRequireDefault(_DrawerShade);

var _DrawerPane = require('./DrawerPane');

var _DrawerPane2 = _interopRequireDefault(_DrawerPane);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var Temporary = function (_Component) {
  _inherits(Temporary, _Component);

  _createClass(Temporary, null, [{
    key: 'handleDrawerClick',
    value: function handleDrawerClick(event) {
      event.stopPropagation();
    }
  }, {
    key: 'isWrongPointer',
    value: function isWrongPointer(pointerType) {
      return pointerType && pointerType !== 'touch';
    }
  }]);

  function Temporary(props) {
    _classCallCheck(this, Temporary);

    var _this = _possibleConstructorReturn(this, (Temporary.__proto__ || Object.getPrototypeOf(Temporary)).call(this, props));

    _this.handleShadeClick = _this.handleShadeClick.bind(_this);
    _this.handleTransitionend = _this.handleTransitionend.bind(_this);
    _this.handleTouchstart = _this.handleTouchstart.bind(_this);
    _this.handleTouchmove = _this.handleTouchmove.bind(_this);
    _this.handleTouchend = _this.handleTouchend.bind(_this);
    _this.close = _this.close.bind(_this);
    _this.state = {};
    return _this;
  }

  _createClass(Temporary, [{
    key: 'componentWillReceiveProps',
    value: function componentWillReceiveProps(newProps) {
      if (this.props.open !== newProps.open) {
        this.setState({ animating: true });
      }
    }
  }, {
    key: 'handleShadeClick',
    value: function handleShadeClick() {
      this.close();
    }
  }, {
    key: 'handleTouchstart',
    value: function handleTouchstart(_ref, drawerWidth) {
      var pointerType = _ref.pointerType,
          touches = _ref.touches,
          pageX = _ref.pageX;

      if (!this.props.open) {
        return;
      }

      if (Temporary.isWrongPointer(pointerType)) {
        return;
      }

      this.startX = touches ? touches[0].pageX : pageX;
      this.touchingSideNav = true;
      this.drawerWidth = drawerWidth;
      this.setState({ currentX: this.startX });
    }
  }, {
    key: 'handleTouchmove',
    value: function handleTouchmove(_ref2) {
      var pointerType = _ref2.pointerType,
          touches = _ref2.touches,
          pageX = _ref2.pageX;

      if (Temporary.isWrongPointer(pointerType)) {
        return;
      }
      var currentX = touches ? touches[0].pageX : pageX;
      this.setState({ currentX: currentX });
    }
  }, {
    key: 'handleTouchend',
    value: function handleTouchend(_ref3) {
      var pointerType = _ref3.pointerType;

      if (Temporary.isWrongPointer(pointerType)) {
        return;
      }
      var newPosition = this.calculateDrawerPosition();
      this.touchingSideNav = false;

      this.setState({ animating: true });

      // Did the user close the drawer by more than 50%?
      if (Math.abs(newPosition / this.drawerWidth) >= 0.5) {
        this.close();
      }
    }
  }, {
    key: 'close',
    value: function close() {
      var onClose = this.props.onClose;

      if (typeof onClose === 'function') {
        onClose();
      }
    }
  }, {
    key: 'handleTransitionend',
    value: function handleTransitionend() {
      this.setState({ animating: false });
    }
  }, {
    key: 'calculateDrawerPosition',
    value: function calculateDrawerPosition() {
      if (this.touchingSideNav) {
        return Math.min(0, this.state.currentX - this.startX);
      }
      return null;
    }
  }, {
    key: 'calculateShadeOpacity',
    value: function calculateShadeOpacity(position) {
      if (this.touchingSideNav) {
        return Math.max(0, 1 + 1 * (position / this.drawerWidth));
      }
      return null;
    }
  }, {
    key: 'render',
    value: function render() {
      var _props = this.props,
          className = _props.className,
          children = _props.children,
          open = _props.open,
          otherProps = _objectWithoutProperties(_props, ['className', 'children', 'open']);

      var animating = this.state.animating;


      var childs = _react2.default.Children.map(children, function (child) {
        return _react2.default.cloneElement(child, { temporary: true });
      });

      var position = this.calculateDrawerPosition();
      var opacity = this.calculateShadeOpacity(position);

      return _react2.default.createElement(
        _DrawerShade2.default,
        _extends({
          animating: animating,
          className: className,
          onClick: this.handleShadeClick,
          onTouchend: this.handleTouchend,
          onTouchmove: this.handleTouchmove,
          opacity: opacity,
          open: open
        }, otherProps),
        _react2.default.createElement(
          _DrawerPane2.default,
          {
            onClick: Temporary.handleDrawerClick,
            onTouchstart: this.handleTouchstart,
            animating: animating,
            onTransitionend: this.handleTransitionend,
            position: position
          },
          childs
        )
      );
    }
  }]);

  return Temporary;
}(_react.Component);

Temporary.propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  open: _propTypes2.default.bool,
  onClose: _propTypes2.default.func,
  header: _propTypes2.default.oneOf([_propTypes2.default.string, _propTypes2.default.node])
};
exports.default = Temporary;